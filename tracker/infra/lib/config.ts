type ConfigKey = "vpc" | "kafka" | "eks";

export type Config = {
  prod: string;
  test: string;
  components: {
    [key in ConfigKey]: {
      name: string;
      props?: any;
    };
  };
};

const config: Config = {
  prod: "ProdStack",
  test: "TestStack",
  components: {
    vpc: {
      name: "TrackerVPC",
      props: {
        privSubnet: "tracker-subnet-priv",
        pubSubnet: "tracker-subnet-pub",
      },
    },
    kafka: {
      name: "BrokerMessage",
    },
    eks: {
      name: "EKSCluster",
      props: {
        nodeGroup: {
          id: "EKSClusterNodeGroup-1",
          intanceType: "t3.nano",
        },
      },
    },
  },
};

export default config;
