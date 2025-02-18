package ingesters

import (
	"time"

	"github.com/deepfence/ThreatMapper/deepfence_utils/directory"
	ingestersUtil "github.com/deepfence/ThreatMapper/deepfence_utils/utils/ingesters"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func CommitFuncCloudCompliance(ns string, data []ingestersUtil.CloudCompliance) error {
	ctx := directory.NewContextWithNameSpace(directory.NamespaceID(ns))
	driver, err := directory.Neo4jClient(ctx)
	if err != nil {
		return err
	}

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	if err != nil {
		return err
	}
	defer session.Close()

	tx, err := session.BeginTransaction(neo4j.WithTxTimeout(30 * time.Second))
	if err != nil {
		return err
	}
	defer tx.Close()

	if _, err = tx.Run(`
		UNWIND $batch as row
		MERGE (n:CloudCompliance{node_id: row.scan_id + "--" + row.control_id + "--" + COALESCE(row.resource, row.account_id, ""), resource:row.resource, scan_id: row.scan_id, control_id: row.control_id})
		MERGE (m:CloudResource{node_id: row.resource})
		SET n+= row
		WITH n, m
		MATCH (l:CloudComplianceScan{node_id: n.scan_id})
		MERGE (l) -[r:DETECTED]-> (n)
		SET r.masked = false`,
		map[string]interface{}{"batch": CloudCompliancesToMaps(data)}); err != nil {
		return err
	}

	return tx.Commit()
}

func CloudCompliancesToMaps(ms []ingestersUtil.CloudCompliance) []map[string]interface{} {
	res := []map[string]interface{}{}
	for _, v := range ms {
		res = append(res, v.ToMap())
	}
	return res
}
