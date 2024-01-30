
/*Expiresフィールドが設定されていない場合、そのクッキーはブラウザを閉じると同時に削除される
 有効期限の指定方法は２通り。
①Expiresフィールドにtime.Time型の値を設定する方法。
②RawExpiresフィールドに秒単位の数値を設定する方法。
 */
 */
 MaxAgeフィールドに秒数を設定する方法。
 */
 */

type Cookie struct {
	Nmae string
	Value string
	Path string
	Domain string
	Expires time.Time
	RawExpires string
	MaxAge int
	Secure bool
	HttpOnly bool
	Raw string
	Unparsed []string
}
