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

/**
 * Get all NFTs for a collection
 * @param {string} contractAddress
 * @param {string | undefined} prevPageKey
 * @returns Nft[]
 */
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

/**
 * Get all owners for a collection
 * @param {string} contractAddress
 * @returns string[]
 */
export async function getOwnerForCollection(contractAddress) {
  const { owners } = await alchemy.nft.getOwnersForContract(contractAddress, {
    withTokenBalances: true,
  });
  return owners;
}

/**
 * Combine the results of RESULT and OWNERS in a way that corresponds to 721.
 * @param {string[]} owners
 * @param {Nft[]} result
 * @returns Nft[]
 */
export async function combineTo721(owners, result) {
  return result.map((nft) => {
    for (const owner of owners) {
      const flat = owner.tokenBalances.map((b) => b.tokenId);
      // In ERC721, there is only one owner, so it is returned in an array of length 1.
      if (flat.includes(nft.id.tokenId)) {
        return {
          ...nft,
          owners: [owner.ownerAddress],
        };
      }
    }
  });
}
