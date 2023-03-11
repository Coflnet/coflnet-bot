package utils

func IsUserMod(roles [] string) bool {
    for _, role := range roles {
        if role == ModRole() {
            return true
        }
    }

    return false
}

func IsUserHelper(roles [] string) bool {
    for _, role := range roles {
        if role == HelperRole() {
            return true
        }
    }

    return false
}
