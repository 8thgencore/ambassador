import axios from "axios";
import { Dispatch, useEffect, useState } from "react";
import { Navigate } from "react-router-dom";
import { User } from "../models/user";
import Menu from "./Menu";
import Nav from "./Nav";
import { connect } from "react-redux";
import { setUser } from "../redux/actions/setUserAction";

const Layout = (props: any) => {
  const [redirect, setRedirect] = useState(false);
  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    (async () => {
      try {
        const { data } = await axios.get("user");

        setUser(data);
        props.setUser(data);
      } catch (e) {
        setRedirect(true);
      }
    })();
  }, []);

  if (redirect) {
    return <Navigate to={"/login"} />;
  }

  return (
    <div>
      <Nav />
      <div className="container-fluid">
        <div className="row">
          <Menu />
          <main className="col-md-9 ms-sm-auto col-lg-10 px-md-4 gy-4">
            <div className="table-responsive">{props.children}</div>
          </main>
        </div>
      </div>
    </div>
  );
};

const mapStateToProps = (state: { user: User }) => ({
  user: state.user,
});

const mapDispatchToProps = (dispatch: Dispatch<any>) => ({
  setUser: (user: User) => dispatch(setUser(user)),
});

export default connect(mapStateToProps, mapDispatchToProps)(Layout);
