package handlers

import (
	"bioskop-app/database"
	"bioskop-app/models"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Endpoint CRUD
// CreateBioskop godoc
// @Summary      Tambah bioskop baru
// @Description  Menambahkan bioskop ke database
// @Tags         bioskop
// @Accept       json
// @Produce      json
// @Param        bioskop body models.Bioskop true "Data bioskop"
// @Success      201 {object} models.Bioskop
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /bioskop [post]
func CreateBioskop(c *gin.Context) {
	var input models.Bioskop
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON tidak valid"})
		return
	}

	if input.Nama == "" || input.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	query := `INSERT INTO bioskop (nama, lokasi, rating) VALUES ($1, $2, $3) RETURNING id`
	err := database.DB.QueryRow(query, input.Nama, input.Lokasi, input.Rating).Scan(&input.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data"})
		return
	}

	c.JSON(http.StatusCreated, input)
}

// GetAllBioskop godoc
// @Summary      Daftar semua bioskop
// @Description  Mengambil semua data bioskop
// @Tags         bioskop
// @Produce      json
// @Success      200 {array} models.Bioskop
// @Failure      500 {object} map[string]string
// @Router       /bioskop [get]
func GetAllBioskop(c *gin.Context) {
	rows, err := database.DB.Query(`SELECT id, nama, lokasi, rating FROM bioskop`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}
	defer rows.Close()

	var bioskops []models.Bioskop
	for rows.Next() {
		var b models.Bioskop
		if err := rows.Scan(&b.ID, &b.Nama, &b.Lokasi, &b.Rating); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca data"})
			return
		}
		bioskops = append(bioskops, b)
	}

	c.JSON(http.StatusOK, bioskops)
}

// GetBioskopByID godoc
// @Summary      Detail bioskop berdasarkan ID
// @Description  Mengambil detail bioskop berdasarkan ID
// @Tags         bioskop
// @Produce      json
// @Param        id path int true "ID bioskop"
// @Success      200 {object} models.Bioskop
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /bioskop/{id} [get]
func GetBioskopByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var b models.Bioskop
	query := `SELECT id, nama, lokasi, rating FROM bioskop WHERE id = $1`
	err = database.DB.QueryRow(query, id).Scan(&b.ID, &b.Nama, &b.Lokasi, &b.Rating)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}

	c.JSON(http.StatusOK, b)
}

// UpdateBioskop godoc
// @Summary      Perbarui data bioskop
// @Description  Memperbarui data bioskop berdasarkan ID
// @Tags         bioskop
// @Accept       json
// @Produce      json
// @Param        id path int true "ID bioskop"
// @Param        bioskop body models.Bioskop true "Data bioskop baru"
// @Success      200 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /bioskop/{id} [put]
func UpdateBioskop(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var input models.Bioskop
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON tidak valid"})
		return
	}

	if input.Nama == "" || input.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	query := `UPDATE bioskop SET nama = $1, lokasi = $2, rating = $3 WHERE id = $4`
	result, err := database.DB.Exec(query, input.Nama, input.Lokasi, input.Rating, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui data"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

// DeleteBioskop godoc
// @Summary      Hapus bioskop
// @Description  Menghapus bioskop berdasarkan ID
// @Tags         bioskop
// @Produce      json
// @Param        id path int true "ID bioskop"
// @Success      200 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /bioskop/{id} [delete]
func DeleteBioskop(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	query := `DELETE FROM bioskop WHERE id = $1`
	result, err := database.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
