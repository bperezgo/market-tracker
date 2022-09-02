type ConfigKey = "vpc" | "kafka";

export type Config = {
  prod: string;
  test: string;
  stacks: {
    [key in ConfigKey]: {
      name: string;
    };
  };
};

const config: Config = {
  prod: "prod",
  test: "test",
  stacks: {
    vpc: {
      name: "TrackerVPC",
    },
    kafka: {
      name: "BrokerMessage",
    },
  },
};

export default config;
