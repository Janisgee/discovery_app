export const fetchPlaceImage = async (place, searchFor, country) => {
  const data = { place_name: place, search_for: searchFor, country: country };
  const request = new Request(
    `${process.env.NEXT_PUBLIC_API_SERVER_BASE}/api/getDisplayPlaceImage`,
    {
      method: "POST", // HTTP method
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify(data),
    },
  );

  try {
    const response = await fetch(request);

    if (response.ok) {
      const imageInfo = await response.json(); // Use json() to handle HTML response

      if (imageInfo) {
        return imageInfo.image_url;
      }
    } else {
      console.error("Error fetching place image details:", response.statusText);
    }
  } catch (error) {
    console.error("Error fetching place image details:", error);
  }
};
