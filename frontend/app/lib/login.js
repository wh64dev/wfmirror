"use server";

export const login = async (currentState, formData) => {
    const raw = {
        username: formData.get("username"),
        password: formData.get("password")
    };

    return "error";
}
