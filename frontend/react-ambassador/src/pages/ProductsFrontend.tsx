import axios from "axios";
import React, { useEffect, useState } from "react";
import Layout from "../components/Layout";
import { Filters } from "../models/filters";
import { Product } from "../models/product";
import Products from "./Products";

const ProductsFrontend = () => {
  const frontendUrl = "products/frontend";
  const [filterProducts, setFilterProducts] = useState<Product[]>([]);
  const [allProducts, setAllProducts] = useState<Product[]>([]);
  const [filters, setFilters] = useState<Filters>({
    q: "",
    sort: "",
    page: 1,
  });
  const [lastPage, setLastPage] = useState(0);

  const perPage = 9;

  useEffect(() => {
    (async () => {
      const { data } = await axios.get(frontendUrl);
      if (data) {
        setAllProducts(data);
        setFilterProducts(data);
        setLastPage(Math.ceil(data.length / perPage));
      }
    })();
  }, []);

  useEffect(() => {
    let products = allProducts.filter(
      (p) =>
        p.title.toLocaleLowerCase().indexOf(filters.q.toLowerCase()) >= 0 ||
        p.description.toLowerCase().indexOf(filters.q.toLowerCase()) >= 0
    );

    if (filters.sort === "asc") {
      products.sort((a: Product, b: Product) => {
        if (a.price > b.price) {
          return 1;
        }
        if (a.price < b.price) {
          return -1;
        }
        return 0;
      });
    } else if (filters.sort === "desc") {
      products.sort((a: Product, b: Product) => {
        if (a.price > b.price) {
          return -1;
        }
        if (a.price < b.price) {
          return 1;
        }
        return 0;
      });
    }

    setLastPage(Math.ceil(products.length / perPage));
    setFilterProducts(products.slice(0, filters.page * perPage));
  }, [filters]);

  return (
    <Layout>
      <Products
        products={filterProducts}
        filters={filters}
        setFilters={setFilters}
        lastPage={lastPage}
      />
    </Layout>
  );
};

export default ProductsFrontend;
