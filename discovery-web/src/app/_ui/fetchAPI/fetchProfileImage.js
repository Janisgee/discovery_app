export const fetchProfileImage = async () => {
  const request = new Request(
    `${process.env.NEXT_PUBLIC_API_SERVER_BASE}/api/displayUserProfileImage`,
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
