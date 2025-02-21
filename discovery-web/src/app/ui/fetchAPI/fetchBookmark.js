export const fetchAllBookmark = async (router) => {
  const request = new Request("http://localhost:8080/api/getAllBookmark", {
    method: "GET", // HTTP method
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
  });

  try {
    const response = await fetch(request);
    if (response.ok) {
      const htmlContent = await response.json(); // Use json() to handle HTML response
      console.log("Received content:", htmlContent);
      return htmlContent;
    } else {
      if (response.status == 401) {
        alert(
          `Please login again as 10 mins session expired without taking action.`,
        );
        router.push(`/login`);
      }
      console.error("Error fetching bookmark place:", response.statusText);
      throw new Error(response.statusText);
    }
  } catch (error) {
    console.error("Error fetching bookmark place:", error);
    return error;
  }
};
