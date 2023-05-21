import NoParticipants from "@/components/noParticipants";
import { useClientsQuery } from "@/graphql";
import { useEffect } from "react";
import { AiOutlineCheck } from "react-icons/ai";

export default function Home() {
  const [{ error, fetching, data }, refetch] = useClientsQuery();

  useEffect(() => {
    const interval = setInterval(() => {
      refetch();
    }, 1500);
    return () => clearInterval(interval);
  }, []);

  if (error) {
    return <div>error</div>;
  }

  if (fetching && !data) {
    return <div>Loading...</div>;
  }

  return (
    <div className="flex flex-row justify-center">
      {data?.clients.length === 0 && <NoParticipants />}
      {data?.clients.map((client) => (
        <div
          key={client.name}
          className="pt-3 bg-gray-300 w-80 text-center flex flex-col justify-center rounded"
        >
          <AiOutlineCheck className="w-10 h-10 text-green-700 self-center" />
          <div className="text-lg text-gray-700">
            {client.name}
            <img src={client.avatarUrl} className="h-20 mx-auto mt-3" />
          </div>
        </div>
      ))}
    </div>
  );
}
