import { useCookies } from "next-client-cookies";

export default function CheckLoginExpire() {
  const cookies = useCookies();
  const checkBoolean = cookies.get("IS_LOGIN");

  const checkLoginExpire = async (checkBoolean) => {
    console.log("hi inside a function", checkBoolean);
    console.log("hi inside a function", checkBoolean.checkBoolean);
    if (checkBoolean.checkBoolean == "1") {
      console.log("Now is expired.. 1");
      // router.push("/login");
    }
  };

  setInterval(checkLoginExpire, 6000, { checkBoolean });

  return <></>;
}
