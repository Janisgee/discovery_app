export const fetchAllBookmark = async () => {
  const request = new Request(
    `${process.env.NEXT_PUBLIC_API_SERVER_BASE}/api/getAllBookmark`,
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
