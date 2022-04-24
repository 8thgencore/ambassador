import Header from "./Header";
import Nav from "./Nav";

const Layout = (props: any) => {
  return (
    <div>
      <Nav />
      <main>
        <Header />
        <div className="album py-5 bg-light">
          <div className="container">{props.children}</div>
        </div>
      </main>
    </div>
  );
};

export default Layout;
