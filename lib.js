import { Network, Alchemy } from "alchemy-sdk";
import { spinner } from "@clack/prompts";
import * as dotenv from "dotenv";

/**
 * @typedef { import("alchemy-sdk").Nft } Nft
 * @typedef { import("alchemy-sdk").NftContractOwner } NftContractOwner
 */

dotenv.config();

const s = spinner();
const settings = {
  apiKey: process.env.ALCHEMY_API_KEY,
  network: Network.ETH_MAINNET,
};
const alchemy = new Alchemy(settings);

export async function getCollectionInfo(contractAddress) {
  s.start("Fetching collection info");
  const res = await alchemy.nft.getContractMetadata(contractAddress);
  s.stop("Done.");

  return res;
}

/**
 * Get all NFTs for a collection
 * @param {string} contractAddress
 * @param {string | undefined} prevToken
 * @returns Promise<Nft[]>
 */
export async function getNftsForCollection(
  contractAddress,
  prevToken,
  prevData = []
) {
  let result = prevData;
  s.start("Fetching 100 NFTs");

  const searchParams = new URLSearchParams({
    contractAddress,
    withMetadata: true,
  });

  if (prevToken != null) {
    searchParams.append("startToken", prevToken);
  }

  const res = await fetch(
    `https://eth-mainnet.g.alchemy.com/nft/v2/${
      settings.apiKey
    }/getNFTsForCollection?${searchParams.toString()}`
  );
  const { nfts, nextToken } = await res.json();

  result = result.concat(nfts);

  s.stop(`Done. Current result: ${result.length}. Next cursor: ${nextToken}`);

  if (nextToken == null) {
    return result;
  }

  return await getNftsForCollection(contractAddress, nextToken, result);
}

/**
 * Get all owners for a collection
 * @param {string} contractAddress
 * @returns Promise<NftContractOwner[]>
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
 * @returns ?[] TODO: Fix type
 */
export async function combineTo721(owners, result) {
  const combined = [];

  for (const nft of result) {
    const holders = [];

    for (const owner of owners) {
      // Assuming ERC721, so I don't see b.balance.
      const flat = owner.tokenBalances.map((b) => b.tokenId);

      if (flat.includes(nft.id.tokenId)) {
        holders.push(owner.ownerAddress);
      }
    }

    combined.push({
      ...nft,
      owners: holders,
    });
  }

  return combined;
}