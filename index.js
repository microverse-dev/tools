import fs from "fs-extra";
import {
  getCollectionInfo,
  getNftsForCollection,
  getOwnerForCollection,
  combineTo721,
} from "./lib.js";
import { intro, outro, text, confirm, cancel } from "@clack/prompts";
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

    fs.removeSync("result");
    fs.mkdirSync("result");
    fs.writeFileSync(
      "result/collection.json",
      // TODO: now supports only ERC721
      JSON.stringify(await combineTo721(owners, result)),
      "utf8"
    );

    outro("Done!🎉 See result directory.");
  } catch (e) {
    console.error(e);
  }
}

main();
