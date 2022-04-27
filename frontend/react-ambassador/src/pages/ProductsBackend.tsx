import axios from "axios";
import React, { useEffect, useState } from "react";
import Layout from "../components/Layout";
import { Filters } from "../models/filters";
import { Product } from "../models/product";
import Products from "./Products";

const ProductsBackend = () => {
  const backendUrl = "products/backend";
  const [products, setProducts] = useState<Product[]>([]);
  const [filters, setFilters] = useState<Filters>({
    q: "",
    sort: "",
    page: 1,
  });
  const [lastPage, setLastPage] = useState(0);

  useEffect(() => {
    (async () => {
      const arr = [];

      if (filters.q) {
        arr.push(`q=${filters.q}`);
      }

      if (filters.sort) {
        arr.push(`sort=${filters.sort}`);
      }

      if (filters.page) {
        arr.push(`page=${filters.page}`);
      }

      const { data } = await axios.get(backendUrl + "?" + arr.join("&"));

      if (data.data) {
        console.log(data);
        setProducts(filters.page === 1 ? data.data : [...products, ...data.data]);
        setLastPage(data.last_page);
      }
    })();
  }, [filters]);

  return (
    <Layout>
      <Products products={products} filters={filters} setFilters={setFilters} lastPage={lastPage} />
    </Layout>
  );
};

export default ProductsBackend;
