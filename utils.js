import fs from "fs-extra";
import { groupBy } from "./lib.js";
import { createObjectCsvWriter } from "csv-writer";

export function clear() {
  fs.removeSync("result");
  fs.mkdirSync("result");
}

function writeFile(filename, data) {
  fs.writeFileSync(filename, data, "utf8");
}

export async function writeCollectionResult(combined) {
  // TODO: now supports only ERC721
  writeFile("result/collection.json", JSON.stringify(combined));
}

export async function writeGroupByTraitResult(trait, combined) {
  writeFile(
    `result/group-by-${trait}.json`,
    JSON.stringify(await groupBy(trait, combined))
  );
}

export async function writeCSV(trait, owners, combined) {
  const grouped = await groupBy(trait, combined);

  // all holders export
  const allHoders = owners.map((owner) => {
    const { ownerAddress, tokenBalances } = owner;
    const balance = tokenBalances.reduce((acc, cur) => acc + cur.balance, 0);

    return {
      address: ownerAddress,
      balance,
    };
  });

  const allHoldersCsvWriter = createObjectCsvWriter({
    path: "result/all-holders.csv",
    header: [
      { id: "address", title: "address" },
      { id: "balance", title: "balance" },
    ],
  });

  await allHoldersCsvWriter.writeRecords(allHoders);

  for await (const key of Object.keys(grouped)) {
    const csvWriter = createObjectCsvWriter({
      path: `result/${key}.csv`,
      header: [{ id: "owner", title: "address" }],
    });

    const records = grouped[key].reduce((acc, cur) => {
      if (cur.owners.length === 0) return;

      acc.push(...cur.owners);
      return acc;
    }, []);

    await csvWriter.writeRecords(records.map((owner) => ({ owner })));
  }
}
