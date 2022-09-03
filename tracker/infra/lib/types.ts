import { Vpc } from "aws-cdk-lib/aws-ec2";

export interface WithVpc {
  vpc: Vpc;
}
