import { VFC } from "react";
import { Route, Switch } from "react-router";
import Login from "../components/pages/Login";
import Singup from "../components/pages/Singup";
import Top from "../components/pages/Top";
import { todoRoutes } from "./TodoRoutes";

export const Router: VFC = () => {
  return(
    <Switch>
      <Route exact path="/">
        <Top/>
      </Route>
      <Route exact path="/login">
        <Login/>
      </Route>
      <Route exact path="/signup">
        <Singup/>
      </Route>
      <Route path="/todo" render={({match: {url}})=> (
        <Switch>
          {todoRoutes.map((route) => (
            <Route 
              key={route.path} 
              exact={route.exact} 
              path={`${url}${route.path}`}
            >
              {route.children}
            </Route>
          ))}
        </Switch>
      )}/>
    </Switch>
  )
}