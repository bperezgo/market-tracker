import { App, Stack } from "aws-cdk-lib";
import { BrokerMessage } from "./kafka-construct";
import { TrackerVpc } from "./vpc-construct";
import config from "./config";
import { K8SCluster } from "./eks-cluster-construct";

type EnvProps = {
  prod: boolean;
};

export class TrackerStack extends Stack {
  constructor(scope: App, id: string, props?: EnvProps) {
    super(scope, id);

    const trackerVpc = new TrackerVpc(this, config.components.vpc.name);

    const kafka = new BrokerMessage(this, config.components.kafka.name, {
      vpc: trackerVpc.vpc,
      securityGroup: trackerVpc.kafkaSecurityGroup,
    });

    const eks = new K8SCluster(this, config.components.eks.name, {
      vpc: trackerVpc.vpc,
    });
  }
}
