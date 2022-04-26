import axios from "axios";
import React, { useEffect, useState } from "react";
import Layout from "../components/Layout";
import { Product } from "../models/product";
import Products from "./Products";

const ProductsFrontend = () => {
  const backendUrl = "products/frontend";
  const [products, setProducts] = useState<Product[]>([]);

  useEffect(() => {
    (async () => {
      const { data } = await axios.get(backendUrl);

      if (data) {
        setProducts(data);
      }
    })();
  });

  return (
    <Layout>
      <Products products={products} />
    </Layout>
  );
};

export default ProductsFrontend;
