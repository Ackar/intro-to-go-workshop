import { Client } from "@/components/client";
import { useLevel2Query } from "@/graphql";
import { useEffect } from "react";

export default function Level2() {
  const [{ fetching, data, error }, refetch] = useLevel2Query({
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
      {data?.level2.map((e) => (
        <div key={e.client.name}>
          <Client client={e.client} />
          <div className="">
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
