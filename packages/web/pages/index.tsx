import { useState } from "react";
import { useFetchInspection, useInspect } from "../lib/api";
import styles from "./index.module.scss";

export function Index() {
  const [url, setUrl] = useState<string>("");
  const [requestId, setRequestId] = useState<string>("");
  const { data } = useFetchInspection(requestId, {
    onError: () => setRequestId(""),
  });
  const { mutate: inspect } = useInspect();

  const onSubmit = () => {
    inspect(url, { onSuccess: (data) => setRequestId(data.id) });
  };

  return (
    <div className={styles.page}>
      <div className="wrapper">
        <div className="container">
          <div id={styles.welcome}>
            <h1>
              <span>Malicious URL detection</span>
              GIMI
            </h1>
          </div>
          <div id={styles.form} className="card shadow">
            <div className={styles.wrapper}>
              <h2>Inspect</h2>
              <label>
                <div>URL</div>
                <input
                  type="text"
                  value={url}
                  onChange={(e) => setUrl(e.target.value)}
                />
              </label>
              <button onClick={onSubmit}>Submit</button>
            </div>
          </div>
          {data?.result?.malicious && (
            <div id={styles.malicious} className="card shadow">
              <h2>This URL looks malicious.</h2>
              <p>These are the reasons.</p>
              <ul>
                {data.result.reasons.map((reason) => (
                  <li key={reason}>{reason}</li>
                ))}
              </ul>
            </div>
          )}
          <div id={styles.result} className="card shadow">
            <div className={styles.wrapper}>
              <h2>Result</h2>
              {requestId && !data?.result && !data?.error && (
                <div>Inspecting...</div>
              )}
              {data?.error && <pre>Error: {data.error}</pre>}
              {data?.result && (
                <table>
                  <tbody>
                    <tr>
                      <th>original url is...</th>
                    </tr>
                    <tr>
                      <td>{data.url}</td>
                    </tr>
                    <tr>
                      <th>redirects you to...</th>
                    </tr>
                    {data.result.locations?.map((location) => (
                      <tr key={location}>
                        <td>{location}</td>
                      </tr>
                    ))}
                    <tr>
                      <th>sends your data to...</th>
                    </tr>
                    {data.result.sendingTo?.map((url) => (
                      <tr key={url}>
                        <td>{url}</td>
                      </tr>
                    ))}
                    <tr>
                      <th>communicates with...</th>
                    </tr>
                    {data.result.hosts?.map((host) => (
                      <tr key={host}>
                        <td>{host}</td>
                      </tr>
                    ))}
                    <tr>
                      <th>requests to...</th>
                    </tr>
                    {data.result.urls?.map((url) => (
                      <tr key={url}>
                        <td>{url}</td>
                      </tr>
                    ))}
                    <tr>
                      <th>looks like...</th>
                    </tr>
                    <tr>
                      <td>
                        <img src={data.result.screenshot} alt="" />
                      </td>
                    </tr>
                  </tbody>
                </table>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Index;
