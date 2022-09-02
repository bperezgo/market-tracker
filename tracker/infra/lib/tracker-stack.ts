import { App, Stack } from "aws-cdk-lib";
import { BrokerMessage } from "./kafka-construct";
import { TrackerVpc } from "./vpc-construct";
import config from "./config";

type EnvProps = {
  prod: boolean;
};

export class TrackerStack extends Stack {
  constructor(scope: App, id: string, props?: EnvProps) {
    super(scope, id);

    const trackerVpc = new TrackerVpc(this, config.stacks.vpc.name);

    const kafka = new BrokerMessage(this, config.stacks.kafka.name, {
      vpc: trackerVpc,
    });
  }
}
