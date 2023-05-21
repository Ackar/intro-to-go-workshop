import { Client } from "@/components/client";
import ClientError from "@/components/clientError";
import NoParticipants from "@/components/noParticipants";
import { useLevel3Query } from "@/graphql";
import { useEffect } from "react";

export default function Level2() {
  const [{ fetching, data, error }, refetch] = useLevel3Query({
    requestPolicy: "cache-and-network",
  });

  useEffect(() => {
    const interval = setInterval(() => {
      refetch();
    }, 1000);
    return () => clearInterval(interval);
  }, []);

  if (fetching && !data) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error :/ {JSON.stringify(error)}</div>;
  }

  return (
    <div>
      {data?.level3.length === 0 && <NoParticipants />}
      {data?.level3.map((e) => (
        <div key={e.client.name}>
          <Client client={e.client} />
          {e.error && <ClientError error={e.error} />}
          <div className="">
            <img src={e.gifUrl} className="h-36" />
          </div>
        </div>
      ))}
    </div>
  );
}
