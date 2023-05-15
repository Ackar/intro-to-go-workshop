import { Client } from "@/components/client";
import { useLevel1Query } from "@/graphql";
import { useEffect } from "react";

export default function Level1() {
  const [{ fetching, data, error }, refetch] = useLevel1Query({
    requestPolicy: "cache-and-network",
  });

  useEffect(() => {
    const interval = setInterval(() => {
      refetch();
    }, 3000);
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
      {data?.level1.map((e) => (
        <div key={e.client.name}>
          <Client client={e.client} />
          <div className="">
            {e.colors.map((c) => (
              <div
                key={c}
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
