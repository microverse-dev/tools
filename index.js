import {
  getCollectionInfo,
  getNftsForCollection,
  getOwnerForCollection,
  combineTo721,
} from "./lib.js";
import {
  clear,
  writeCollectionResult,
  writeGroupByTraitResult,
  writeCSV,
} from "./utils.js";
import {
  intro,
  outro,
  text,
  confirm,
  cancel,
  select,
  spinner,
} from "@clack/prompts";
import { exit } from "node:process";
import { setTimeout } from "timers/promises";

async function main() {
  intro("NFT Collection snapshot");

  const contractAddress = await text({
    message: "What is collection's contract address?",
    placeholder: "0x...",
    validate: (value) => {
      if (value.length === 0) return "ContractAddress is required";
      if (!value.startsWith("0x")) return "ContractAddress must start with 0x";
    },
  });

  const contractInfo = await getCollectionInfo(contractAddress);
  const checked = await confirm({
    message: `Is the collection name ${contractInfo.name}?`,
  });

  if (!checked) {
    cancel("Canceled.");
    exit(0);
  }

  try {
    const owners = await getOwnerForCollection(contractAddress);
    const result = await getNftsForCollection(contractAddress);

    if (result == null) {
      outro("No NFTs found.");
      return;
    }

    const combined = await combineTo721(owners, result);

    clear();

    const trait = await select({
      message: "What key do you use to group?",
      options: result[0].metadata.attributes.map((a) => ({
        value: a.trait_type,
        label: a.trait_type,
      })),
    });

    const s = spinner();
    s.start("Exporting...");

    await Promise.all([
      writeCollectionResult(combined),
      writeGroupByTraitResult(trait, combined),
      writeCSV(trait, owners, combined),
    ]);
    s.stop();

    await setTimeout(1000);

    outro("Done!🎉 See result directory.");
  } catch (e) {
    console.error(e);
  }
}

main();
