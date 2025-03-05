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
    } else {
      console.error("Error fetching search country:", response.statusText);
      throw new Error(response.statusText);
    }
  } catch (error) {
    console.error("Error fetching user new profile image:", error);
    return error;
  }
};
