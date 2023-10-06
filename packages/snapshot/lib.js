import { Network, Alchemy } from "alchemy-sdk";
import { spinner } from "@clack/prompts";

/**
 * @typedef { import("alchemy-sdk").Nft } Nft
 * @typedef { import("alchemy-sdk").NftContractOwner } NftContractOwner
 */

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

  // We want metadata information, so we use the REST API.
  // ref: https://docs.alchemy.com/reference/getnftsforcollection
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
export async function combine(owners, result) {
  const combined = [];

  for (const nft of result) {
    const holders = [];

    for (const owner of owners) {
      for (const tokenBalance of owner.tokenBalances) {
        if (tokenBalance.tokenId === nft.id.tokenId) {
          holders.push({
            owner: owner.ownerAddress,
            balance: tokenBalance.balance,
          });
        }
      }
    }

    combined.push({
      ...nft,
      owners: holders,
    });
  }

  return combined;
}

/**
 * Group the result by a key
 * @param {string} trait
 * @param {any[]} result
 */
export async function groupBy(trait, combined) {
  return combined.reduce((acc, nft) => {
    const attribute = nft.metadata.attributes.find(
      (a) => a.trait_type === trait
    );
    const key = attribute?.value ?? "Unknown";

    if (!acc[key]) {
      acc[key] = [];
    }

    acc[key].push(nft);

    return acc;
  }, {});
}
