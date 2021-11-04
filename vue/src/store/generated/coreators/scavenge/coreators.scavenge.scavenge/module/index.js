// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgSubmitScavenge } from "./types/scavenge/tx";
import { MsgCommitSolution } from "./types/scavenge/tx";
import { MsgRevealSolution } from "./types/scavenge/tx";
const types = [
    ["/coreators.scavenge.scavenge.MsgSubmitScavenge", MsgSubmitScavenge],
    ["/coreators.scavenge.scavenge.MsgCommitSolution", MsgCommitSolution],
    ["/coreators.scavenge.scavenge.MsgRevealSolution", MsgRevealSolution],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgSubmitScavenge: (data) => ({ typeUrl: "/coreators.scavenge.scavenge.MsgSubmitScavenge", value: data }),
        msgCommitSolution: (data) => ({ typeUrl: "/coreators.scavenge.scavenge.MsgCommitSolution", value: data }),
        msgRevealSolution: (data) => ({ typeUrl: "/coreators.scavenge.scavenge.MsgRevealSolution", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
