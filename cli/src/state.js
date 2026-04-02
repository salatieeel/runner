import { RUNNER_STATE_FILE } from "./constants.js";
import { readJson, writeJson } from "./utils.js";

const DEFAULT_STATE = {
  assinador: {},
  simulador: {},
  downloads: {}
};

export async function loadState() {
  return readJson(RUNNER_STATE_FILE, structuredClone(DEFAULT_STATE));
}

export async function saveState(state) {
  await writeJson(RUNNER_STATE_FILE, state);
}

export async function updateState(mutator) {
  const state = await loadState();
  await mutator(state);
  await saveState(state);
  return state;
}
