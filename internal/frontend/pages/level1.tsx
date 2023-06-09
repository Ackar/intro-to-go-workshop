import { Client } from "@/components/client";
import ClientError from "@/components/clientError";
import NoParticipants from "@/components/noParticipants";
import { useLevel1Query } from "@/graphql";
import { useEffect } from "react";

export default function Level1() {
  const [{ fetching, data, error }, refetch] = useLevel1Query({
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
      {data?.level1.length === 0 && <NoParticipants />}
      {data?.level1.map((e) => (
        <div key={e.client.name}>
          <Client client={e.client} />
          <div className="">
            {e.error && <ClientError error={e.error} />}
            {e.colors.map((c, idx) => (
              <div
                key={idx}
                className="w-10 h-10 inline-block mx-0.5 shadow"
                style={{
                  backgroundColor: c,
                }}
              ></div>
            ))}
          </div>
        </div>
      ))}
    </div>
  );
}
