import axios from "axios";
import React, { useEffect, useState } from "react";
import Layout from "../components/Layout";
import { Product } from "../models/product";
import Products from "./Products";

const ProductsBackend = () => {
    const backendUrl = "products/backend";
    const [products, setProducts] = useState<Product[]>([]);
  
    useEffect(() => {
      (async () => {
        const { data } = await axios.get(backendUrl);
  
        if (data.data) {
          console.log(data);
          setProducts(data.data);
        }
      })();
    });
  
    return (
      <Layout>
        <Products products={products} />
      </Layout>
    );
  };

export default ProductsBackend;
