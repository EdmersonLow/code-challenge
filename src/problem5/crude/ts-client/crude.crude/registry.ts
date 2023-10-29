import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgDeleteResource } from "./types/crude/tx";
import { MsgCreateResource } from "./types/crude/tx";
import { MsgUpdateResource } from "./types/crude/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/crude.crude.MsgDeleteResource", MsgDeleteResource],
    ["/crude.crude.MsgCreateResource", MsgCreateResource],
    ["/crude.crude.MsgUpdateResource", MsgUpdateResource],
    
];

export { msgTypes }