export const fetchInputControl = async (searchWord) => {
  const request = new Request(
    `http://localhost:8080/api/place/autocomplete?search=${encodeURIComponent(searchWord)}`,
    {
      method: "GET", // HTTP method
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
    },
  );

  try {
    const response = await fetch(request);
    if (response.ok) {
      const htmlContent = await response.json(); // Use json() to handle HTML response
      console.log("Received content:", htmlContent);
      return htmlContent;
    } else {
      console.error("Error fetching bookmark place:", response.statusText);
      throw new Error(response.statusText);
    }
  } catch (error) {
    console.error("Error fetching bookmark place:", error);
    return error;
  }
};
