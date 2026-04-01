import os from "node:os";
import path from "node:path";

export const APP_NAME = "runner";
export const DEFAULT_ASSINADOR_PORT = 18080;
export const DEFAULT_SIMULADOR_PORT = 19090;
export const DEFAULT_IDLE_TIMEOUT_MINUTES = 15;
export const JDK_VERSION = "21";
export const RUNNER_HOME = path.join(os.homedir(), ".hubsaude");
export const RUNNER_STATE_FILE = path.join(RUNNER_HOME, "runner-state.json");
export const RUNNER_CACHE_DIR = path.join(RUNNER_HOME, "cache");
export const RUNNER_TOOLS_DIR = path.join(RUNNER_HOME, "tools");
export const RUNNER_LOG_DIR = path.join(RUNNER_HOME, "logs");
export const ASSINADOR_JAR_CANDIDATES = [
  path.resolve(process.cwd(), "assinador", "build", "assinador.jar"),
  path.resolve(process.cwd(), "dist", "assinador.jar"),
  path.resolve(process.cwd(), "assinador.jar")
];
export const DEFAULT_SIMULADOR_REPO = "kyriosdata/simulador-hubsaude";
