import { ClientFragment } from "@/graphql";

export function Client({ client }: { client: ClientFragment }) {
  return (
    <div className="flex flex-row bg-gray-300 w-max px-3 py-1 rounded my-2">
      <img src={client.avatarUrl} className="max-w-10 h-10 mr-3" />
      <div className="self-center">{client.name}</div>
    </div>
  );
}
