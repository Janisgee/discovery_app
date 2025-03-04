"use client";

import { Button } from "@/app/_ui/buttons";

export default function App() {
  return (
    <div className="p-20 text-center">
      <div className="text-right">
        <Button useFor="Sign Up" link="/signup" color="btn-grey" />
      </div>
      <div className="my-20 ">
        <h1>Discover Your Side!</h1>
      </div>
      <Button useFor="Login" link="/login" color="btn-violet" />
    </div>
  );
}
