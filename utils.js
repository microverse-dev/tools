import fs from "fs-extra";
import { groupBy } from "./lib.js";

export function clear() {
  fs.removeSync("result");
  fs.mkdirSync("result");
}

export async function writeCollectionResult(combined) {
  fs.writeFileSync(
    "result/collection.json",
    // TODO: now supports only ERC721
    JSON.stringify(combined),
    "utf8"
  );
}

export async function writeGroupByTraitResult(trait, combined) {
  fs.writeFileSync(
    `result/group-by-${trait}.json`,
    JSON.stringify(await groupBy(trait, combined)),
    "utf8"
  );
}
