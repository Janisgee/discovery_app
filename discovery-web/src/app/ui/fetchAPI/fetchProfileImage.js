export const fetchProfileImage = async () => {
  const request = new Request(
    "http://localhost:8080/api/displayUserProfileImage",
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
      const htmlContent = await response.json();
      return htmlContent;
    }
    if (!response.ok) {
      throw new Error(`Failed to fetch: ${response.statusText}`);
    }
  } catch (error) {
    console.error("Error fetching user new profile image:", error);
  }
};
