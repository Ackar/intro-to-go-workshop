import { MdErrorOutline } from "react-icons/md";

export default function ClientError({ error }: { error: string }) {
  return (
    <div className="bg-red-100 flex flex-row px-3 py-1 rounded w-max">
      <MdErrorOutline className="self-center mr-2" />
      Error: {error}
    </div>
  );
}
