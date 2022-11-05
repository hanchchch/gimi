import styles from "./signin.module.scss";

export function Signin() {
  return (
    <div className={styles.page}>
      <div className="wrapper">
        <div className="container">
          <div id={styles.welcome}>
            <h1>
              <span> Hello there, </span>
              Welcome 👋
            </h1>
          </div>
          <div id={styles.hero} className="card">
            <h2>
              <svg
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth="2"
                  d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z"
                />
              </svg>
              <span>This website is completely fine.</span>
            </h2>
          </div>
          <div id={styles.form} className="card shadow">
            <div className={styles.wrapper}>
              <h2>Sign In</h2>
              <form>
                <label>
                  <span>Username</span>
                  <input type="text" />
                </label>
                <label>
                  <span>Password</span>
                  <input type="password" />
                </label>
                <button type="submit">Submit</button>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Signin;
