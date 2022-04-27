import { useEffect } from "react";
import { useRouter } from "next/dist/client/router";
import Layout from "../components/Layout";
import axios from "axios";
import constants from "../constants";

const Success = () => {
  const router = useRouter();
  const { source } = router.query;
  useEffect(() => {
    if (source !== undefined) {
      (async () => {
        await axios.post(`${constants.endpoint}/orders/confirm`, {
          source,
        });
      })();
    }
  });
  return (
    <Layout>
      <div className="py-5 text-center">
        <h2>Success</h2>
        <p className="lead">Your purchase has been completed!</p>
      </div>
    </Layout>
  );
};

export default Success;
