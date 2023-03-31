import { Network, Alchemy } from "alchemy-sdk";
import { spinner } from "@clack/prompts";
import * as dotenv from "dotenv";
import { setTimeout } from "timers/promises";

dotenv.config();

const s = spinner();
const settings = {
  apiKey: process.env.ALCHEMY_API_KEY,
  network: Network.ETH_MAINNET,
};
const alchemy = new Alchemy(settings);
const result = [];

export async function getNftsForCollection(contractAddress, prevPageKey) {
  s.start("Fetching 100 NFTs");
  const { nfts, pageKey } = await alchemy.nft.getNftsForContract(
    contractAddress,
    { pageKey: prevPageKey, pageSize: 100 }
  );

  result.push(...nfts);

  s.stop(`Done. Current result: ${result.length}. Next cursor: ${pageKey}`);

  if (pageKey == null) {
    return result;
  }

  await setTimeout(500);
  await getNftsForCollection(contractAddress, pageKey);
}

export async function getOwnerForCollection(contractAddress) {
  const { owners } = await alchemy.nft.getOwnersForContract(contractAddress);
  return owners;
}
