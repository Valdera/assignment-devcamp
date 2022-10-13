package product

const (
	addProductQuery = `
	INSERT INTO products (
		name,
		price,
		description,
		variant,
		discount,
		created_at,
		updated_at,
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
	) returning id
`
	getProductByIdQuery = `
	SELECT
		id,
		name,
		price,
		description,
		variant,
		discount,
		created_at,
		updated_at,
	FROM
		products
	WHERE
		id=$1
`

	getProductAllQuery = `
	SELECT
		*
	FROM
		products
`

	updateProductQuery = `
	UPDATE
		products
	SET
		name=$1,
		price=$2,
		description=$3,
		variant=$4,
		discount=$5,
		created_at=$6,
		updated_at=$7
	WHERE
		id=$8
	returning id	
`

	deleteProductQuery = `
	DELETE
	FROM products
	WHERE id=$1
`
)
