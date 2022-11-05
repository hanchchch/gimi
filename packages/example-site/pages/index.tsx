export const getServerSideProps = async () => {
  return {
    redirect: {
      destination: "/signin",
    },
  };
};

export function Index() {
  return <div></div>;
}

export default Index;
