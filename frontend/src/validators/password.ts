const passwordRule = (password: string) => {
    if(false === /\d/.test(password))
        return false;

    return true;
}

export default passwordRule;