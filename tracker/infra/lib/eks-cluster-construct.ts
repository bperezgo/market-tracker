import * as eks from "aws-cdk-lib/aws-eks";
import * as ec2 from "aws-cdk-lib/aws-ec2";
import { Construct } from "constructs";
import config from "./config";
import { WithVpc } from "./types";

export interface K8SClusterProps extends WithVpc {}

export class K8SCluster extends Construct {
  constructor(scope: Construct, id: string, props: K8SClusterProps) {
    super(scope, id);

    const cluster = new eks.Cluster(this, config.components.eks.name, {
      version: eks.KubernetesVersion.V1_21,
      vpc: props.vpc,
      defaultCapacity: 0,
    });

    cluster.addNodegroupCapacity(config.components.eks.props.nodeGroup.id, {
      instanceTypes: [
        new ec2.InstanceType(config.components.eks.props.nodeGroup.intanceType),
      ],
      minSize: 2,
    });
  }
}
