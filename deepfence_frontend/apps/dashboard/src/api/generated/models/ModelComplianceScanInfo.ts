/* tslint:disable */
/* eslint-disable */
/**
 * Deepfence ThreatMapper
 * Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.
 *
 * The version of the OpenAPI document: 2.0.0
 * Contact: community@deepfence.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ModelComplianceScanInfo
 */
export interface ModelComplianceScanInfo {
    /**
     * 
     * @type {Array<string>}
     * @memberof ModelComplianceScanInfo
     */
    benchmark_types: Array<string> | null;
    /**
     * 
     * @type {string}
     * @memberof ModelComplianceScanInfo
     */
    node_id: string;
    /**
     * 
     * @type {string}
     * @memberof ModelComplianceScanInfo
     */
    node_type: string;
    /**
     * 
     * @type {string}
     * @memberof ModelComplianceScanInfo
     */
    scan_id: string;
    /**
     * 
     * @type {{ [key: string]: number; }}
     * @memberof ModelComplianceScanInfo
     */
    severity_counts: { [key: string]: number; } | null;
    /**
     * 
     * @type {string}
     * @memberof ModelComplianceScanInfo
     */
    status: string;
    /**
     * 
     * @type {number}
     * @memberof ModelComplianceScanInfo
     */
    updated_at: number;
}

/**
 * Check if a given object implements the ModelComplianceScanInfo interface.
 */
export function instanceOfModelComplianceScanInfo(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "benchmark_types" in value;
    isInstance = isInstance && "node_id" in value;
    isInstance = isInstance && "node_type" in value;
    isInstance = isInstance && "scan_id" in value;
    isInstance = isInstance && "severity_counts" in value;
    isInstance = isInstance && "status" in value;
    isInstance = isInstance && "updated_at" in value;

    return isInstance;
}

export function ModelComplianceScanInfoFromJSON(json: any): ModelComplianceScanInfo {
    return ModelComplianceScanInfoFromJSONTyped(json, false);
}

export function ModelComplianceScanInfoFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelComplianceScanInfo {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'benchmark_types': json['benchmark_types'],
        'node_id': json['node_id'],
        'node_type': json['node_type'],
        'scan_id': json['scan_id'],
        'severity_counts': json['severity_counts'],
        'status': json['status'],
        'updated_at': json['updated_at'],
    };
}

export function ModelComplianceScanInfoToJSON(value?: ModelComplianceScanInfo | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'benchmark_types': value.benchmark_types,
        'node_id': value.node_id,
        'node_type': value.node_type,
        'scan_id': value.scan_id,
        'severity_counts': value.severity_counts,
        'status': value.status,
        'updated_at': value.updated_at,
    };
}

