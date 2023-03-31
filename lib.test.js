import { combineTo721 } from "./lib.js";

const owners = [
  {
    ownerAddress: "0x24941b659992908A74AE345551aD07F1D6E5Fd84",
    tokenBalances: [
      {
        tokenId:
          "0x0000000000000000000000000000000000000000000000000000000000000001",
        balance: 1,
      },
    ],
  },
  {
    ownerAddress: "0x680a48D0311d1363C6b97eD0dC97A78b65B35539",
    tokenBalances: [
      {
        tokenId:
          "0x0000000000000000000000000000000000000000000000000000000000000002",
        balance: 1,
      },
      {
        tokenId:
          "0x0000000000000000000000000000000000000000000000000000000000000003",
        balance: 1,
      },
    ],
  },
];
const result = [
  {
    contract: {
      address: "0x0bda9d7185a9885ecb1770d4793389be5da2e576",
    },
    id: {
      tokenId:
        "0x0000000000000000000000000000000000000000000000000000000000000001",
      tokenMetadata: {
        tokenType: "ERC721",
      },
    },
    title: "IROIRO #1",
    description:
      "IROIRO is a colorful NFT collectible designed by artist sashimi & CyberZ, and IROIRO is uwuLabs Family project.You can choose and get the NFT in any color you like.",
    tokenUri: {
      gateway:
        "https://alchemy.mypinata.cloud/ipfs/QmeDy2yHzjsr3Nr1MZ3qfCWupD2Lg9rZo29y4jdnhhTdKm/1",
      raw: "ipfs://QmeDy2yHzjsr3Nr1MZ3qfCWupD2Lg9rZo29y4jdnhhTdKm/1",
    },
    media: [
      {
        gateway:
          "https://nft-cdn.alchemy.com/eth-mainnet/dba9c3b5cdf7c4da6112ea6a9630939d",
        thumbnail:
          "https://res.cloudinary.com/alchemyapi/image/upload/thumbnailv2/eth-mainnet/dba9c3b5cdf7c4da6112ea6a9630939d",
        raw: "ipfs://QmSNxvZBxpqhk3363N5Bqnt5JkQwtR6k8WLH78x9Yegk6c/1.png",
        format: "png",
        bytes: 1182360,
      },
    ],
    metadata: {
      name: "IROIRO #1",
      description:
        "IROIRO is a colorful NFT collectible designed by artist sashimi & CyberZ, and IROIRO is uwuLabs Family project.You can choose and get the NFT in any color you like.",
      edition: "1",
      image: "ipfs://QmSNxvZBxpqhk3363N5Bqnt5JkQwtR6k8WLH78x9Yegk6c/1.png",
      external_url: "https://iroiro.world/",
      attributes: [
        {
          value: "Manga",
          trait_type: "Background",
        },
        {
          value: "None",
          trait_type: "Halo",
        },
        {
          value: "Hane short",
          trait_type: "Hair",
        },
        {
          value: "Sakura",
          trait_type: "Body",
        },
        {
          value: "None",
          trait_type: "Mole",
        },
        {
          value: "Pokan",
          trait_type: "Mouth",
        },
        {
          value: "kyorokyoro red",
          trait_type: "Eye",
        },
        {
          value: "Heart",
          trait_type: "Eyedeco",
        },
        {
          value: "Tsuri",
          trait_type: "Eyebrow",
        },
        {
          value: "Star",
          trait_type: "Facedeco",
        },
        {
          value: "Heart Sakura",
          trait_type: "Hand",
        },
        {
          value: "Shirt",
          trait_type: "Dress",
        },
        {
          value: "Tsuntsun",
          trait_type: "Bang",
        },
        {
          value: "Strawberry",
          trait_type: "Accessory",
        },
        {
          value: "None",
          trait_type: "Glasses",
        },
        {
          value: "AKA",
          trait_type: "color",
        },
      ],
    },
    timeLastUpdated: "2023-03-07T02:10:11.656Z",
    contractMetadata: {
      name: "IROIRO",
      symbol: "IROIRO",
      tokenType: "ERC721",
      contractDeployer: "0x7d475ca71eef776998717d1b6bc8aa69576d8810",
      deployedBlockNumber: 16654125,
      openSea: {
        floorPrice: 0.1049,
        collectionName: "IROIRO",
        safelistRequestStatus: "approved",
        imageUrl:
          "https://i.seadn.io/gcs/files/259e052f055ca1b75dfee782a87d1a6a.png?w=500&auto=format",
        description:
          "IROIRO is a colorful NFT collectible designed by artist [sashimi](https://foundation.app/@sashimi_a_u) & [CyberZ](https://cyber-z.co.jp/en), and IROIRO is [uwu Labs](https://uwulabs.com/) Family project. \nYou can choose and get the NFT in any color you like????",
        externalUrl: "https://iroiro.world/",
        twitterUsername: "IROIRO_NFT",
        discordUrl: "https://discord.gg/iroiro",
        lastIngestedAt: "2023-03-18T06:40:53.000Z",
      },
    },
  },
  {
    contract: {
      address: "0x0bda9d7185a9885ecb1770d4793389be5da2e576",
    },
    id: {
      tokenId:
        "0x0000000000000000000000000000000000000000000000000000000000000002",
      tokenMetadata: {
        tokenType: "ERC721",
      },
    },
    title: "IROIRO #2",
    description:
      "IROIRO is a colorful NFT collectible designed by artist sashimi & CyberZ, and IROIRO is uwuLabs Family project.You can choose and get the NFT in any color you like.",
    tokenUri: {
      gateway:
        "https://alchemy.mypinata.cloud/ipfs/QmeDy2yHzjsr3Nr1MZ3qfCWupD2Lg9rZo29y4jdnhhTdKm/2",
      raw: "ipfs://QmeDy2yHzjsr3Nr1MZ3qfCWupD2Lg9rZo29y4jdnhhTdKm/2",
    },
    media: [
      {
        gateway:
          "https://nft-cdn.alchemy.com/eth-mainnet/ec4037d00f15645973f4c1c0478c0f38",
        thumbnail:
          "https://res.cloudinary.com/alchemyapi/image/upload/thumbnailv2/eth-mainnet/ec4037d00f15645973f4c1c0478c0f38",
        raw: "ipfs://QmSNxvZBxpqhk3363N5Bqnt5JkQwtR6k8WLH78x9Yegk6c/2.png",
        format: "png",
        bytes: 1003567,
      },
    ],
    metadata: {
      name: "IROIRO #2",
      description:
        "IROIRO is a colorful NFT collectible designed by artist sashimi & CyberZ, and IROIRO is uwuLabs Family project.You can choose and get the NFT in any color you like.",
      edition: "2",
      image: "ipfs://QmSNxvZBxpqhk3363N5Bqnt5JkQwtR6k8WLH78x9Yegk6c/2.png",
      external_url: "https://iroiro.world/",
      attributes: [
        {
          value: "Effect",
          trait_type: "Background",
        },
        {
          value: "None",
          trait_type: "Halo",
        },
        {
          value: "Short",
          trait_type: "Hair",
        },
        {
          value: "Yuki",
          trait_type: "Body",
        },
        {
          value: "None",
          trait_type: "Mole",
        },
        {
          value: "Pokan",
          trait_type: "Mouth",
        },
        {
          value: "shitoshito purple",
          trait_type: "Eye",
        },
        {
          value: "Heart",
          trait_type: "Eyedeco",
        },
        {
          value: "Komari",
          trait_type: "Eyebrow",
        },
        {
          value: "Star",
          trait_type: "Facedeco",
        },
        {
          value: "Fox Yuki",
          trait_type: "Hand",
        },
        {
          value: "Miko",
          trait_type: "Dress",
        },
        {
          value: "Normal",
          trait_type: "Bang",
        },
        {
          value: "Fox",
          trait_type: "Accessory",
        },
        {
          value: "None",
          trait_type: "Glasses",
        },
        {
          value: "AO",
          trait_type: "color",
        },
      ],
    },
    timeLastUpdated: "2023-03-07T07:47:35.207Z",
    contractMetadata: {
      name: "IROIRO",
      symbol: "IROIRO",
      tokenType: "ERC721",
      contractDeployer: "0x7d475ca71eef776998717d1b6bc8aa69576d8810",
      deployedBlockNumber: 16654125,
      openSea: {
        floorPrice: 0.1049,
        collectionName: "IROIRO",
        safelistRequestStatus: "approved",
        imageUrl:
          "https://i.seadn.io/gcs/files/259e052f055ca1b75dfee782a87d1a6a.png?w=500&auto=format",
        description:
          "IROIRO is a colorful NFT collectible designed by artist [sashimi](https://foundation.app/@sashimi_a_u) & [CyberZ](https://cyber-z.co.jp/en), and IROIRO is [uwu Labs](https://uwulabs.com/) Family project. \nYou can choose and get the NFT in any color you like????",
        externalUrl: "https://iroiro.world/",
        twitterUsername: "IROIRO_NFT",
        discordUrl: "https://discord.gg/iroiro",
        lastIngestedAt: "2023-03-18T06:40:53.000Z",
      },
    },
  },
  {
    contract: {
      address: "0x0bda9d7185a9885ecb1770d4793389be5da2e576",
    },
    id: {
      tokenId:
        "0x0000000000000000000000000000000000000000000000000000000000000003",
      tokenMetadata: {
        tokenType: "ERC721",
      },
    },
    title: "IROIRO #3",
    description:
      "IROIRO is a colorful NFT collectible designed by artist sashimi & CyberZ, and IROIRO is uwuLabs Family project.You can choose and get the NFT in any color you like.",
    tokenUri: {
      gateway:
        "https://alchemy.mypinata.cloud/ipfs/QmeDy2yHzjsr3Nr1MZ3qfCWupD2Lg9rZo29y4jdnhhTdKm/1",
      raw: "ipfs://QmeDy2yHzjsr3Nr1MZ3qfCWupD2Lg9rZo29y4jdnhhTdKm/1",
    },
    media: [
      {
        gateway:
          "https://nft-cdn.alchemy.com/eth-mainnet/dba9c3b5cdf7c4da6112ea6a9630939d",
        thumbnail:
          "https://res.cloudinary.com/alchemyapi/image/upload/thumbnailv2/eth-mainnet/dba9c3b5cdf7c4da6112ea6a9630939d",
        raw: "ipfs://QmSNxvZBxpqhk3363N5Bqnt5JkQwtR6k8WLH78x9Yegk6c/1.png",
        format: "png",
        bytes: 1182360,
      },
    ],
    metadata: {
      name: "IROIRO #1",
      description:
        "IROIRO is a colorful NFT collectible designed by artist sashimi & CyberZ, and IROIRO is uwuLabs Family project.You can choose and get the NFT in any color you like.",
      edition: "1",
      image: "ipfs://QmSNxvZBxpqhk3363N5Bqnt5JkQwtR6k8WLH78x9Yegk6c/1.png",
      external_url: "https://iroiro.world/",
      attributes: [
        {
          value: "Manga",
          trait_type: "Background",
        },
        {
          value: "None",
          trait_type: "Halo",
        },
        {
          value: "Hane short",
          trait_type: "Hair",
        },
        {
          value: "Sakura",
          trait_type: "Body",
        },
        {
          value: "None",
          trait_type: "Mole",
        },
        {
          value: "Pokan",
          trait_type: "Mouth",
        },
        {
          value: "kyorokyoro red",
          trait_type: "Eye",
        },
        {
          value: "Heart",
          trait_type: "Eyedeco",
        },
        {
          value: "Tsuri",
          trait_type: "Eyebrow",
        },
        {
          value: "Star",
          trait_type: "Facedeco",
        },
        {
          value: "Heart Sakura",
          trait_type: "Hand",
        },
        {
          value: "Shirt",
          trait_type: "Dress",
        },
        {
          value: "Tsuntsun",
          trait_type: "Bang",
        },
        {
          value: "Strawberry",
          trait_type: "Accessory",
        },
        {
          value: "None",
          trait_type: "Glasses",
        },
        {
          value: "SHIRO",
          trait_type: "color",
        },
      ],
    },
    timeLastUpdated: "2023-03-07T02:10:11.656Z",
    contractMetadata: {
      name: "IROIRO",
      symbol: "IROIRO",
      tokenType: "ERC721",
      contractDeployer: "0x7d475ca71eef776998717d1b6bc8aa69576d8810",
      deployedBlockNumber: 16654125,
      openSea: {
        floorPrice: 0.1049,
        collectionName: "IROIRO",
        safelistRequestStatus: "approved",
        imageUrl:
          "https://i.seadn.io/gcs/files/259e052f055ca1b75dfee782a87d1a6a.png?w=500&auto=format",
        description:
          "IROIRO is a colorful NFT collectible designed by artist [sashimi](https://foundation.app/@sashimi_a_u) & [CyberZ](https://cyber-z.co.jp/en), and IROIRO is [uwu Labs](https://uwulabs.com/) Family project. \nYou can choose and get the NFT in any color you like????",
        externalUrl: "https://iroiro.world/",
        twitterUsername: "IROIRO_NFT",
        discordUrl: "https://discord.gg/iroiro",
        lastIngestedAt: "2023-03-18T06:40:53.000Z",
      },
    },
  },
];

describe("lib", () => {
  test("combineTo721", async () => {
    const expected = [
      {
        contract: {
          address: "0x0bda9d7185a9885ecb1770d4793389be5da2e576",
        },
        id: {
          tokenId:
            "0x0000000000000000000000000000000000000000000000000000000000000001",
          tokenMetadata: {
            tokenType: "ERC721",
          },
        },
        title: "IROIRO #1",
        description:
          "IROIRO is a colorful NFT collectible designed by artist sashimi & CyberZ, and IROIRO is uwuLabs Family project.You can choose and get the NFT in any color you like.",
        tokenUri: {
          gateway:
            "https://alchemy.mypinata.cloud/ipfs/QmeDy2yHzjsr3Nr1MZ3qfCWupD2Lg9rZo29y4jdnhhTdKm/1",
          raw: "ipfs://QmeDy2yHzjsr3Nr1MZ3qfCWupD2Lg9rZo29y4jdnhhTdKm/1",
        },
        media: [
          {
            gateway:
              "https://nft-cdn.alchemy.com/eth-mainnet/dba9c3b5cdf7c4da6112ea6a9630939d",
            thumbnail:
              "https://res.cloudinary.com/alchemyapi/image/upload/thumbnailv2/eth-mainnet/dba9c3b5cdf7c4da6112ea6a9630939d",
            raw: "ipfs://QmSNxvZBxpqhk3363N5Bqnt5JkQwtR6k8WLH78x9Yegk6c/1.png",
            format: "png",
            bytes: 1182360,
          },
        ],
        metadata: {
          name: "IROIRO #1",
          description:
            "IROIRO is a colorful NFT collectible designed by artist sashimi & CyberZ, and IROIRO is uwuLabs Family project.You can choose and get the NFT in any color you like.",
          edition: "1",
          image: "ipfs://QmSNxvZBxpqhk3363N5Bqnt5JkQwtR6k8WLH78x9Yegk6c/1.png",
          external_url: "https://iroiro.world/",
          attributes: [
            {
              value: "Manga",
              trait_type: "Background",
            },
            {
              value: "None",
              trait_type: "Halo",
            },
            {
              value: "Hane short",
              trait_type: "Hair",
            },
            {
              value: "Sakura",
              trait_type: "Body",
            },
            {
              value: "None",
              trait_type: "Mole",
            },
            {
              value: "Pokan",
              trait_type: "Mouth",
            },
            {
              value: "kyorokyoro red",
              trait_type: "Eye",
            },
            {
              value: "Heart",
              trait_type: "Eyedeco",
            },
            {
              value: "Tsuri",
              trait_type: "Eyebrow",
            },
            {
              value: "Star",
              trait_type: "Facedeco",
            },
            {
              value: "Heart Sakura",
              trait_type: "Hand",
            },
            {
              value: "Shirt",
              trait_type: "Dress",
            },
            {
              value: "Tsuntsun",
              trait_type: "Bang",
            },
            {
              value: "Strawberry",
              trait_type: "Accessory",
            },
            {
              value: "None",
              trait_type: "Glasses",
            },
            {
              value: "AKA",
              trait_type: "color",
            },
          ],
        },
        owners: ["0x24941b659992908A74AE345551aD07F1D6E5Fd84"],
        timeLastUpdated: "2023-03-07T02:10:11.656Z",
        contractMetadata: {
          name: "IROIRO",
          symbol: "IROIRO",
          tokenType: "ERC721",
          contractDeployer: "0x7d475ca71eef776998717d1b6bc8aa69576d8810",
          deployedBlockNumber: 16654125,
          openSea: {
            floorPrice: 0.1049,
            collectionName: "IROIRO",
            safelistRequestStatus: "approved",
            imageUrl:
              "https://i.seadn.io/gcs/files/259e052f055ca1b75dfee782a87d1a6a.png?w=500&auto=format",
            description:
              "IROIRO is a colorful NFT collectible designed by artist [sashimi](https://foundation.app/@sashimi_a_u) & [CyberZ](https://cyber-z.co.jp/en), and IROIRO is [uwu Labs](https://uwulabs.com/) Family project. \nYou can choose and get the NFT in any color you like????",
            externalUrl: "https://iroiro.world/",
            twitterUsername: "IROIRO_NFT",
            discordUrl: "https://discord.gg/iroiro",
            lastIngestedAt: "2023-03-18T06:40:53.000Z",
          },
        },
      },
      {
        contract: {
          address: "0x0bda9d7185a9885ecb1770d4793389be5da2e576",
        },
        id: {
          tokenId:
            "0x0000000000000000000000000000000000000000000000000000000000000002",
          tokenMetadata: {
            tokenType: "ERC721",
          },
        },
        title: "IROIRO #2",
        description:
          "IROIRO is a colorful NFT collectible designed by artist sashimi & CyberZ, and IROIRO is uwuLabs Family project.You can choose and get the NFT in any color you like.",
        tokenUri: {
          gateway:
            "https://alchemy.mypinata.cloud/ipfs/QmeDy2yHzjsr3Nr1MZ3qfCWupD2Lg9rZo29y4jdnhhTdKm/2",
          raw: "ipfs://QmeDy2yHzjsr3Nr1MZ3qfCWupD2Lg9rZo29y4jdnhhTdKm/2",
        },
        media: [
          {
            gateway:
              "https://nft-cdn.alchemy.com/eth-mainnet/ec4037d00f15645973f4c1c0478c0f38",
            thumbnail:
              "https://res.cloudinary.com/alchemyapi/image/upload/thumbnailv2/eth-mainnet/ec4037d00f15645973f4c1c0478c0f38",
            raw: "ipfs://QmSNxvZBxpqhk3363N5Bqnt5JkQwtR6k8WLH78x9Yegk6c/2.png",
            format: "png",
            bytes: 1003567,
          },
        ],
        metadata: {
          name: "IROIRO #2",
          description:
            "IROIRO is a colorful NFT collectible designed by artist sashimi & CyberZ, and IROIRO is uwuLabs Family project.You can choose and get the NFT in any color you like.",
          edition: "2",
          image: "ipfs://QmSNxvZBxpqhk3363N5Bqnt5JkQwtR6k8WLH78x9Yegk6c/2.png",
          external_url: "https://iroiro.world/",
          attributes: [
            {
              value: "Effect",
              trait_type: "Background",
            },
            {
              value: "None",
              trait_type: "Halo",
            },
            {
              value: "Short",
              trait_type: "Hair",
            },
            {
              value: "Yuki",
              trait_type: "Body",
            },
            {
              value: "None",
              trait_type: "Mole",
            },
            {
              value: "Pokan",
              trait_type: "Mouth",
            },
            {
              value: "shitoshito purple",
              trait_type: "Eye",
            },
            {
              value: "Heart",
              trait_type: "Eyedeco",
            },
            {
              value: "Komari",
              trait_type: "Eyebrow",
            },
            {
              value: "Star",
              trait_type: "Facedeco",
            },
            {
              value: "Fox Yuki",
              trait_type: "Hand",
            },
            {
              value: "Miko",
              trait_type: "Dress",
            },
            {
              value: "Normal",
              trait_type: "Bang",
            },
            {
              value: "Fox",
              trait_type: "Accessory",
            },
            {
              value: "None",
              trait_type: "Glasses",
            },
            {
              value: "AO",
              trait_type: "color",
            },
          ],
        },
        owners: ["0x680a48D0311d1363C6b97eD0dC97A78b65B35539"],
        timeLastUpdated: "2023-03-07T07:47:35.207Z",
        contractMetadata: {
          name: "IROIRO",
          symbol: "IROIRO",
          tokenType: "ERC721",
          contractDeployer: "0x7d475ca71eef776998717d1b6bc8aa69576d8810",
          deployedBlockNumber: 16654125,
          openSea: {
            floorPrice: 0.1049,
            collectionName: "IROIRO",
            safelistRequestStatus: "approved",
            imageUrl:
              "https://i.seadn.io/gcs/files/259e052f055ca1b75dfee782a87d1a6a.png?w=500&auto=format",
            description:
              "IROIRO is a colorful NFT collectible designed by artist [sashimi](https://foundation.app/@sashimi_a_u) & [CyberZ](https://cyber-z.co.jp/en), and IROIRO is [uwu Labs](https://uwulabs.com/) Family project. \nYou can choose and get the NFT in any color you like????",
            externalUrl: "https://iroiro.world/",
            twitterUsername: "IROIRO_NFT",
            discordUrl: "https://discord.gg/iroiro",
            lastIngestedAt: "2023-03-18T06:40:53.000Z",
          },
        },
      },
      {
        contract: {
          address: "0x0bda9d7185a9885ecb1770d4793389be5da2e576",
        },
        id: {
          tokenId:
            "0x0000000000000000000000000000000000000000000000000000000000000003",
          tokenMetadata: {
            tokenType: "ERC721",
          },
        },
        title: "IROIRO #3",
        description:
          "IROIRO is a colorful NFT collectible designed by artist sashimi & CyberZ, and IROIRO is uwuLabs Family project.You can choose and get the NFT in any color you like.",
        tokenUri: {
          gateway:
            "https://alchemy.mypinata.cloud/ipfs/QmeDy2yHzjsr3Nr1MZ3qfCWupD2Lg9rZo29y4jdnhhTdKm/1",
          raw: "ipfs://QmeDy2yHzjsr3Nr1MZ3qfCWupD2Lg9rZo29y4jdnhhTdKm/1",
        },
        media: [
          {
            gateway:
              "https://nft-cdn.alchemy.com/eth-mainnet/dba9c3b5cdf7c4da6112ea6a9630939d",
            thumbnail:
              "https://res.cloudinary.com/alchemyapi/image/upload/thumbnailv2/eth-mainnet/dba9c3b5cdf7c4da6112ea6a9630939d",
            raw: "ipfs://QmSNxvZBxpqhk3363N5Bqnt5JkQwtR6k8WLH78x9Yegk6c/1.png",
            format: "png",
            bytes: 1182360,
          },
        ],
        metadata: {
          name: "IROIRO #1",
          description:
            "IROIRO is a colorful NFT collectible designed by artist sashimi & CyberZ, and IROIRO is uwuLabs Family project.You can choose and get the NFT in any color you like.",
          edition: "1",
          image: "ipfs://QmSNxvZBxpqhk3363N5Bqnt5JkQwtR6k8WLH78x9Yegk6c/1.png",
          external_url: "https://iroiro.world/",
          attributes: [
            {
              value: "Manga",
              trait_type: "Background",
            },
            {
              value: "None",
              trait_type: "Halo",
            },
            {
              value: "Hane short",
              trait_type: "Hair",
            },
            {
              value: "Sakura",
              trait_type: "Body",
            },
            {
              value: "None",
              trait_type: "Mole",
            },
            {
              value: "Pokan",
              trait_type: "Mouth",
            },
            {
              value: "kyorokyoro red",
              trait_type: "Eye",
            },
            {
              value: "Heart",
              trait_type: "Eyedeco",
            },
            {
              value: "Tsuri",
              trait_type: "Eyebrow",
            },
            {
              value: "Star",
              trait_type: "Facedeco",
            },
            {
              value: "Heart Sakura",
              trait_type: "Hand",
            },
            {
              value: "Shirt",
              trait_type: "Dress",
            },
            {
              value: "Tsuntsun",
              trait_type: "Bang",
            },
            {
              value: "Strawberry",
              trait_type: "Accessory",
            },
            {
              value: "None",
              trait_type: "Glasses",
            },
            {
              value: "SHIRO",
              trait_type: "color",
            },
          ],
        },
        owners: ["0x680a48D0311d1363C6b97eD0dC97A78b65B35539"],
        timeLastUpdated: "2023-03-07T02:10:11.656Z",
        contractMetadata: {
          name: "IROIRO",
          symbol: "IROIRO",
          tokenType: "ERC721",
          contractDeployer: "0x7d475ca71eef776998717d1b6bc8aa69576d8810",
          deployedBlockNumber: 16654125,
          openSea: {
            floorPrice: 0.1049,
            collectionName: "IROIRO",
            safelistRequestStatus: "approved",
            imageUrl:
              "https://i.seadn.io/gcs/files/259e052f055ca1b75dfee782a87d1a6a.png?w=500&auto=format",
            description:
              "IROIRO is a colorful NFT collectible designed by artist [sashimi](https://foundation.app/@sashimi_a_u) & [CyberZ](https://cyber-z.co.jp/en), and IROIRO is [uwu Labs](https://uwulabs.com/) Family project. \nYou can choose and get the NFT in any color you like????",
            externalUrl: "https://iroiro.world/",
            twitterUsername: "IROIRO_NFT",
            discordUrl: "https://discord.gg/iroiro",
            lastIngestedAt: "2023-03-18T06:40:53.000Z",
          },
        },
      },
    ];

    expect(await combineTo721(owners, result)).toEqual(expected);
  });
});
