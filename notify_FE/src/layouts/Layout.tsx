import React, { ComponentElement } from "react";
import {
  Container,
  // Sidebar,
  Content,
  // Nav,
  // Sidenav,
  // Icon,
  // Dropdown,
  Header,
  // Navbar,
  Footer,
  // Grid,
  // Row,
  // Col,
  // Avatar,
} from "rsuite";
import { adminRoute } from "./../routes";

import { BrowserRouter as Router, Route, Switch } from "react-router-dom";

import AdminSidebar from "../components/sidebar/AdminSidebar";
import Home from "../view/home";
interface RouteProps123 {
  redirect: string;
  layout: string;
  path: string;
  icon: string;
  name: string;
  component: any;
}
const componentPath = adminRoute.map((prop: any, index: number) => {
  return <Route exactpath={prop.path}>{prop.component}</Route>;
});

export default function Layout(props: any) {
  return (
    <Container className="mainContainer">
      <Router>
        {/* <AdminSidebar routes={adminRoute} {...props} /> */}
        <Container>
          <Header className="header">
            <p>notify</p>
            {/* <Avatar >RS</Avatar> */}
          </Header>
          <Content className="content">
            <Switch>{componentPath}</Switch>
          </Content>
          <Footer className="footer">
            <p>
              Tools team WSO<sub>2</sub> &nbsp;
            </p>
          </Footer>
        </Container>
      </Router>
    </Container>
  );
}
