import AppTemplate from "@/app/ui/template/appTemplate";
import CatagoryTemplate from "@/app/ui/template/catagoryTemplate";

export default function Attraction() {
  return (
    <div>
      <AppTemplate>
        <CatagoryTemplate catagory="attraction"></CatagoryTemplate>
      </AppTemplate>
    </div>
  );
}
