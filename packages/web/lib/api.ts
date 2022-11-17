import { useMutation, useQuery } from "react-query";
import { Inspection } from "./inerface";

const API_HOST = "http://localhost:3000";

export const useFetchInspection = (
  requestId: string,
  options: {
    onError?: (err: Error) => void;
  } = {}
) =>
  useQuery(
    `/inspections/${requestId}`,
    async (): Promise<Inspection> =>
      (await fetch(`${API_HOST}/inspections/${requestId}`)).json(),
    {
      enabled: !!requestId,
      refetchInterval(data, query) {
        return !data?.result ? 1000 : false;
      },
      ...options,
    }
  );

export const useInspect = () =>
  useMutation(
    `/inspections`,
    async (url: string): Promise<Inspection> =>
      (
        await fetch(`${API_HOST}/inspections`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ url }),
        })
      ).json(),
    {}
  );
