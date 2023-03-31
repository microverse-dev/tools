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
} from "./utils.js";
import { intro, outro, text, confirm, cancel, select } from "@clack/prompts";
import { exit } from "node:process";

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
    await writeCollectionResult(combined);

    const useGroupByTrait = await confirm({
      message: "Do you want to group by TRAIT?",
    });

    if (useGroupByTrait) {
      const trait = await select({
        message: "What key do you use to group?",
        options: result[0].metadata.attributes.map((a) => ({
          value: a.trait_type,
          label: a.trait_type,
        })),
      });

      await writeGroupByTraitResult(trait, combined);
    }

    outro("Done!🎉 See result directory.");
  } catch (e) {
    console.error(e);
  }
}

main();
