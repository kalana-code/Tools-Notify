import React, { Component } from 'react';


import  LoginLayOut  from "./../layouts/Login.jsx";
// import User from "layouts/User.jsx"
import { ProtectedRoute } from  "./privateRoute"
import { BrowserRouter, Route, Switch, Redirect } from "react-router-dom";
import auth from './auth'
import Layout from './../layouts/Layout';

class RouteManager extends Component {
    state = {  }
    render() { 
        return ( <BrowserRouter>
            <Switch>
                {/* <Route path="/Register" render={props => <RegisterLayOut {...props}/>} /> */}
                <ProtectedRoute allowRoles={['ADMIN']} exact path="/admin/*" component={Layout} redirectPath="/login"  /> 
                <Route path="/login"  render={props => <LoginLayOut {...props}/>} />
                { auth.getRole() === "ADMIN" ?<Redirect to='/admin/dashboard'  />:<Redirect to='/login'  />}
            </Switch>
          </BrowserRouter> );
    }
}
 
export default RouteManager;