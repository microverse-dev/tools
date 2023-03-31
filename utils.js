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
      header: [
        { id: "owner", title: "address" },
        { id: "balance", title: "balance" },
      ],
    });

    const groupedBalance = grouped[key].reduce((acc, cur) => {
      if (cur.owners.length === 0) return;

      const { owners } = cur;

      for (const owner of owners) {
        const { owner: ownerAddress, balance } = owner;

        acc[ownerAddress] = acc[ownerAddress]
          ? acc[ownerAddress] + balance
          : balance;
      }

      return acc;
    }, {});

    const records = Object.keys(groupedBalance).map((o) => ({
      owner: o,
      balance: groupedBalance[o],
    }));

    await csvWriter.writeRecords(records);
  }
}
