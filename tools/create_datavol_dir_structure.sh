echo "Data volume directory (do not add a slash at the end): "
read

mkdir -vp $REPLY/software/seafile/data
mkdir -vp $REPLY/software/seafile/mysql
